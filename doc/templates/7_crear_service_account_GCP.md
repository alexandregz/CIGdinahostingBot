# Crear conta de servizo en GCP

Para empregar de *backend* Google Cloud Platform (GCP) no *CIGdhBot*, cómpre identificarse empregando algún dos xeitos que provee a plataforma de Google.

O máis doado, mantendo un balance axeitado entre comodidade e seguridade, é empregar unha **conta de servizo** que teña acceso aos ficheiros e directorios do Google Drive que empreguemos.

A continuación describo, máis ou menos pormenorizadamente, o proceso de creación desta e da chave privada a empregar como autenticación para poder acceder aos recursos compartidos por esta conta de servizo.


1. Imos á URL [https://console.cloud.google.com/apis/credentials](https://console.cloud.google.com/apis/credentials)
2. Se non estamos dados de alta, forzarásenos a darnos de alta en Google Cloud.
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/001_alta_GCP.png)
3. Damos de alta o proxecto xenérico
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/002_alta_proxecto.png)
4. Damos de alta a conta de servizo
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/004_alta_credencias_conta_servizo.png)
5. Conta de servizo creada
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/005_creada_conta_servizo.png)
6. Engadimos permisos de *Editor* á conta
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/006_permisos_conta_servizo.png)
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/007_permisos_editor.png)

7. Conta de servizo creada
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/008_permisos_editor_listado.png)
8. Agora creamos a chave privada para autenticarnos na conta de servizo
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/009_crear_keys.png)
9. De tipo JSON
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/010_chave_privada_JSON.png)
10. Chave privada creada
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/011_descarga_chave_privada.png)
11. Descarga no noso equipo
![Alta en GCP se non o estamos xa](../img/conta_servizo_GCP/012_descargada.png)

     
 
