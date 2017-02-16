<?php

 
 
require_once './libs/xmldom.class.php';

$filename   = $_POST['filename'];

$xml        = xml::exportXML($filename);

die($xml);
?>
