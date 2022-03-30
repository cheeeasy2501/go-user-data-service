Есть сервис user-data который хранит информацию о пользователе 
                
    type User struct { 
        Id uint64 
        Email string 
        Password string
        FirstName string
        LastName string 
        Active bool 
    }

В user-data можно отправлять запрос для создания пользователя и обновления информации о пользователе Так же можно
получать пользователя по id или email

Есть второй сервис auth который занимается авторизацией пользователя Все данные о пользователе он запрашивает у сервиса
user-data через gRPC