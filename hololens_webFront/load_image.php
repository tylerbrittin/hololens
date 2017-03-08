<?php


require_once './libs/base64.class.php';

//Model File
$filename   = $_POST['filename'];
//Texture file
$txtr       = $_POST['textureFilename'];

$base64M        = base64::exportBase64($filename);
$base64T        = base64::exportBase64($txtr);

die(json_encode(['Model'=>$base64M, 'Texture'=>$base64T]));
?>
