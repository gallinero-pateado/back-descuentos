# back-descuentos 
Este back contiene el scraping a descuentos de distintos locales de comida, los cuales estan en archivos json que luego se unifican. Tiene un autenticador que verifica que este logeado el usuario, si no esta logeado no accedera a la informacion de descuentos.

Para correr el programa debe ejecutarse el codigo debemos tener en cuenta 2 cosas: 
1- que en la capeta config contenga las credenciales necesarias:
    -serviceAccountKey.json
    -app.env
2- para correr el programa se debe configurar el thunderclient de la siguente manera:
    - Se debe hacer un GET a la siguiente URL "http://localhost:8080/descuentos".

    - En los headers debe existir "Authorization" que debe tener un "Bearer" al inicio y luego el token que entrega el login del usuario.

    - Despues de esto se guardan los cambios y se envia, el resultado debe entregar un JSON estructurado de la siguiente manera.
    
        "message":  "Acceso a descuentos",
        "user_id":  uid,
        "products": allProducts,

Si se hace "go run ." no se dara el resultado esperado, dara {"error":"No se proporcionó token"}