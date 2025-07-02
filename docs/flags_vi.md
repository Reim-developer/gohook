# Cách sử dụng GoHook.

## Lệnh
`wh-send`
* **Mô tả:** Gửi nội dung đã được thiết lập trước trong `settings.toml` hoặc bất kì file cấu hình TOML nào sẵn có miễn là đúng format.

| Tùy chọn     | Kiểu | Mặc định | Mô tả                                                        |
| ------------ | ---- | -------- | ------------------------------------------------------------ |
| `--verbose`  | bool | false    | Hiển thị JSON payload sau khi gửi thành công.                |
| `--dry-run`  | bool | false    | Chế độ chạy thử, chỉ hiển thị JSON payload, không gửi.       |
| `--loop`     | int  | `1`      | Gửi webhook nhiều lần.                                       |
| `--delay`    | int  | `2`      | Chờ giữa các lần gửi, tính bằng giây.                        |
| `--threads`  | int  | `1`      | Tạm chưa sử dụng được ở v1.0.0.                              |
| `--explicit` | bool | false    | Hiển thị thêm các thông tin cần thiết như ID tin nhắn, kênh. |

## Ví dụ:
**Gửi webhook 1 lần duy nhất qua tệp cấu hình:**
```bash
gohook wh-send settings.toml
```

**Chế độ dry, không gửi, chỉ hiển thị JSON payload**
```bash
gohook wh-send settings.toml --dry-run
```

**Gửi 5 lần với 3 giây chờ mỗi lần:**
```bash
gohook wh-send settings.toml --loop 5 --delay 3
```

**Nếu bạn cần mẫu cấu hình TOML, hãy xem ở:**
* https://github.com/Reim-developer/gohook/tree/stable/examples