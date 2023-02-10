# Basic web server

## 필요한 지식

1. 네트워크 관련 컴퓨터 지식

## http pakage

go의 http pakage는 http.ListenAndServe() 메서드를 이용해서 서버 주소를 설정할 수 있다.

요청을 받을 때는 http.HandleFunc() 메서드를 이용해서 요청을 받을 수 있다. http.HandleFunc()은 두 개의 인수를 받는다.

먼저, 패턴을 입력해야 한다. 패턴은 http.ListenAndServe에서 인수로 입력한 주소 다음에 오는 문자열을 의미한다. 예를 들어서 ListneAndServe의 인수에 localhost:8080 이라고 적었다고 가정했을 때, HandleFunc의 패턴으로 "/user" 을 입력했다면,<http://localhost:8080/user> 요청을 의미하는 것이다.

다음으로는 handler 함수를 등록해준다. http.HandleFunc()가 받는 handler 함수는 func(http.ResponseWriter \*http.Request) 타입이어야 한다. http.ResponseWriter는 인터페이스이다. 이 인터페이스를 이용해서 Header를 설정하고 응답을 보낼 수 있다. http.Request를 이용해서는 요청정보를 읽을 수 있다.
