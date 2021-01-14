# arduino_web_control
Управление состоянием реле через интернет

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/web_control.gif)

инфраструктура состоит из:

- сервер со статическим IP (GO + PostgreSQL)
- клиент arduino ( периодически посылает GET запрос о состоянии реле, обратно получает json с состоянием реле и согласно этому json устанавливает состояние реле)
- управление состоянием реле осуществляется через браузер

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/web_control.jpg)




серверная часть написана на Go и включает в себя Postgres сервер для хранения имени, пароля и состояния реле

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/psql.jpeg)

