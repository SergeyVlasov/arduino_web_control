# arduino_web_control
Управление состоянием реле через интернет

- 
![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/web_control.gif)

- 

Для сборки arduino клиента потребуется:

- arduino

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/uno.jpg)

- плата ethernet (w5100)

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/w5100.jpg)


Arduino и плата ethernet собираются в один блок

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/uno%2Bethernet.jpg)


- блок реле

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/relay.jpg)


в итоге схема такая


![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/sceme.jpg)



инфраструктура состоит из:

- сервер со статическим IP (GO + PostgreSQL)
- клиент arduino ( периодически посылает GET запрос о состоянии реле, обратно получает json с состоянием реле и согласно этому json устанавливает состояние реле)
- управление состоянием реле осуществляется через браузер

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/web_control.jpg)




серверная часть написана на Go и включает в себя Postgres сервер для хранения имени, пароля и состояния реле

![til](https://github.com/SergeyVlasov/arduino_web_control/blob/main/media/psql.jpeg)

