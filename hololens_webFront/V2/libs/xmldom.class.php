<?php


 
 

define('UPLOAD_DIR',	getcwd().'/uploads/');

class xml{

    static function mime_content_type($filename) {
    $result = new finfo();

    if (is_resource($result) === true) {
        return $result->file($filename, FILEINFO_MIME_TYPE);
    }

    return false;
}

    static function exportXML($filename){

        //Let's convert the image to base64 first
        $imgData    = static::imageTo64($filename);
        die($imgData);

        //Let's create our XML file
        //Starting by our main DOM
        $dom        = new DOMDocument('1.0');

        //Creating the <Image> node
        $image      = $dom->appendChild($dom->createElement('Image'));

        //Store the name of the image. Create the <Name> node first
        $name       = $image->appendChild($dom->createElement('Name'));
        $name->appendChild($dom->createTextNode($filename));

        //Create and store the image base64 content
        $data       = $image->appendChild($dom->createElement('Data'));
        $data->appendChild($dom->createTextNode($imgData));

        //Generating the XML data
        $dom->formatOutput   = true;

        //Saving the XML as a string to be retuned
        $xml        = $dom->saveXML();

        return $xml;
    }

    static private function imageTo64($filename){

        //Get the absolute path of the image
        $fileDestination    = UPLOAD_DIR.$filename;

        //Let's convert it to base64
        $fileData = base64_encode(file_get_contents($fileDestination));
        //Making sure, this will be recognized as an image
        $data = 'data: '.mime_content_type($fileDestination).';charset=utf8;base64,'.$fileData;

        return $data;
    }
}

?>
