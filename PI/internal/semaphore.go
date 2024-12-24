type Semaphore struct {
C chan struct{}
}

func (s *Semaphore) Acquire() {
s.C <- struct{}{}
}

func (s *Semaphore) Release() {
<-s.C
}

type resultWithError struct {
User users.User
Err  error
}

func DeactivateUsers(usrs []users.User, gCount int) ([]users.User, error) {
// создаем семафор и передаем ему канал с размером буфера равным ограничению на количество одновременно выполняемых горутин
sem := Semaphore{
C: make(chan struct{}, gCount),
}

// канал для сбора результатов
outputCh := make(chan resultWithError, len(usrs))
// канал для оповещения горутин, что мы больше не ждем их выполнения
sgnlCh := make(chan struct{})

output := make([]users.User, 0, len(usrs))

for _, v := range usrs {
go func(usr users.User) {
sem.Acquire()
defer sem.Release()

err := usr.Deactivate()

// если ловим закрытие сигнального канала, то завершаем функцию
select {
case outputCh <- resultWithError{
User: usr,
Err:  err,
}:
case <-sgnlCh:
return
}
}(v)
}

// ждем и собираем результаты
// либо мы получим все, либо выйдет ошибка, по которой мы перестанем читать
for i := len(usrs); i > 0; i-- {
res := <-outputCh
if res.Err != nil {
close(sgnlCh)
return nil, fmt.Errorf("an error occurred: %w", res.Err)
}

output = append(output, res.User)
}

return output, nil
}