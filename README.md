# winter_project2024
## テーブル設計

```marmaid
erDiagram
    users {
        id int
        name str
    }
    rooms {
        id int
        author_id int
        password str
    }
    room_member {
        id int
        room_id int
        user_id int
    }
    compare_image {
        id int
        image_name str
        image str
    }
    result{
		    id int
		    member_id int
		    score float
    }
		user_face{
				id int
				user_id int
				image str
		}


    rooms |o--|| users : "author_id"
    users }o--|| room_member : "user_id"
    users ||--o| user_face : "user_id"
    rooms ||--o{ room_member : "room_id"
    room_member ||--|| result : "member_id"

```
