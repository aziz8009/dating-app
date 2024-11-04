
## API Reference

#### Get all items

```http
  POST /signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your API key |
| `email`     | `string` | **Required**. Your email |
| `password` | `string` | **Required**. Your password |

#### Signup

```http
  POST /login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. Email |
| `password`      | `string` | **Required**. password |

#### Login


```http
  POST /swipe
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `profile_id`      | `number` | **Required**. profile_id |
| `action`      | `string` | **Required**. action |

#### swipe


```http
  Get /swipes/daily
```

#### get swipe daily

```http
  GET /packages
```
#### packages


```http
  POST /purchase
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `package_id`      | `number` | **Required**. package_id |

#### purchase


