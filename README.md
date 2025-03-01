# posts_comments_service
Система для добавления и чтения постов и комментариев с использованием GraphQL.

### Характеристики системы постов:
- Можно просмотреть список постов.
- Можно просмотреть пост и комментарии под ним.
- Пользователь, написавший пост, может запретить оставление комментариев к своему посту.

### Характеристики системы комментариев к постам:
- Комментарии организованы иерархически, позволяя вложенность без ограничений.
- Длина текста комментария ограничена до, например, 2000 символов.
- Система пагинации для получения списка комментариев.

## Запуск с docker:
1) Клонирование репозитория;
2) `make docker.run`

## Запуск тестов:
`make tests.run`

## Остановка:
Остановка и удаление контейнеров - `make docker.down`

## Api

Api представлено в файлах формата .graphqls - [graph](./graph).
Запросы разделены на две категории: посты и комментарии.

Для работы с GraphQL была использована библиотека [gqlgen](https://gqlgen.com/)

## Функционал
### Queries:
- Получение списка постов с поддержкой пагинации (необходимы номер и размер страницы).
- Получение детальной информации о посте, включая комментарии с неограниченной вложенностью и пагинацией.

*Для решения проблемы N+1 комментарии к посту и ответы на них загружаются только в случае, если они запрашиваются в GraphQL запросе.

### Mutations:
- Создание поста ( При создании поста можно указать, разрешены ли комментарии к нему )
- Создание комментария (  Комментарий создается для конкретного поста и может быть ответом на другой комментарий. При добавлении комментария проверяется, разрешено ли оставлять комментарии к данному посту )

### Subscriptions:
- Реализована Подписка на комментарии к определенному посту.

## Выбор хранилища.
Параметр `USE_IN_MEMORY` в [docker-compose.yml](./docker-compose.yml) 11 строка.
`true` - in-memory
`false` - postgresql

## Примеры запросов.
### Создание поста:
```graphql
mutation CreatePost {
    CreatePost(
        post: {
            name: "today's news"
            content: "blah-blah"
            author: "redactor1"
            commentsAllowed: true
        }
    ) {
        id
        createdAt
        name
        author
        content
    }
}
```

### Получение детальной информации о посте:
```graphql
query GetPostById {
    GetPostById(id: 1) {
        id
        createdAt
        name
        author
        content
        commentsAllowed
        comments(page: 1, pageSize: 3) {
            id
            createdAt
            author
            content
            post
            replies {
                id
                createdAt
                author
                content
                post
                replyTo
            }
        }
    }
}
```

### Получение списка постов:
```graphql
query GetAllPosts {
    GetAllPosts(page: 1, pageSize: 3) {
        id
        createdAt
        name
        author
        content
    }
}
```

### Создание комментария:
```graphql
mutation CreateComment {
    CreateComment(input: { author: "author1", content: "interesting.", post: "1" }) {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
```

### Подписка на комментарии поста:
```graphql
subscription CommentsSubscription {
    CommentsSubscription(postId: "1") {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
