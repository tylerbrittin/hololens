    jQuery(document).ready(function() {

        //This is triggered when we click on the submit button
        jQuery("#sendAjax").click(function(){

            // API URL
            var API_URL = 'http://40.87.66.169:5073/additem';

            //Retrieve the image's xml and add it to the serialized array, and send the whole form's POST
            var filename = jQuery("#imageName").val();
            jQuery.post('/load_image.php', {'filename': filename}, function(response){
                //alert(response);
                //Set the input hidden field to hold the image data
                jQuery('input[name=model]').val('<xml>'+response+'</xml>');

                var serialized = jQuery('form[name=submitForm]').serializeFormJSON();


                jQuery.ajax({
                                url: API_URL,
                                type: 'POST',
                                crossDomain: true,
                                data: serialized, //JSON.stringify(data),
                                dataType: 'json',
                                contentType: "application/json",
                                success: function (response) {
                                    var resp = JSON.parse(response)
                                    alert(resp.status);
                                },
                                error: function (xhr, status) {
                                    alert("error");
                                }
                            });


                //Now, let's make our API call
                /*jQuery.post(API_URL, serialized, function(APIResponse){

                    var answer = JSON.parse(APIResponse);
                    alert(answer.status);

                    return false;
                }, 'json');*/

                console.log(serialized);
                return false;
            });
        });

        //File Upload for the image banner
        jQuery('.fileinput-button').click(function(){
            jQuery('#fileupload').click();
        });

        /*//File Upload for the image banner
        jQuery('.fileinput-button').click(function () {
            jQuery('#fileupload').click();
        });*/

        //'use strict';
        // Change this to the location of your server-side upload handler:
        var url = './upload_image.php';

        jQuery('#fileupload').fileupload({
            url: url,
            dataType: 'json',
            maxNumberOfFiles: 1,
            done: function (e, data) {
                jQuery.each(data.result.files, function (index, file) {

                    //Save the uploaded image file name to be used for later
                    image_src = file.name;

                    jQuery("#imageName").val(image_src);
                    /*
                    //To avoid altering the original files, then we do it from here
                    jQuery.post('/<?=ADMIN_PATH.' / '.CMS::getInstance()->module?>/scale_article_banner', {
                        no_html: 'yes',
                        image_name: file.name
                    }, function (returned_data) {

                        //Show the set_banner_form
                        jQuery('.set_banner_form').show();
                        jQuery('#preview_banner').html('<img src="' + file.url + '" class="crop_banner">');

                    });*/
                });
            },
            progressall: function (e, data) {
                var progress = parseInt(data.loaded / data.total * 100, 10);
                jQuery('#progress .progress-bar').css(
                    'width',
                    progress + '%'
                );
            }
        }).prop('disabled', !jQuery.support.fileInput)
            .parent().addClass(jQuery.support.fileInput ? undefined : 'disabled');
    });