function updateUsername(){
    jQuery('#modUsername').val(jQuery('input[name=seller]').val());
    jQuery('#texUsername').val(jQuery('input[name=seller]').val());
}

jQuery(document).ready(function() {

    /*jQuery('#submitDummy').click(function() {
        var API_URL = 'http://40.121.206.106:5073/additem';
        var apiData = JSON.stringify({Category: "furniture", Item: "example lamp", ItemDesc: "itworks", Seller: "Mr. Lamp", Email: "newemail@newemail.com", Price: "$5", Model: "lamp_mod_username", Texture: "something_new"});

        //JSON.stringify(apiData);
        console.log(apiData);

        jQuery.post(API_URL, apiData, function(data){console.log(data); return false;});

        jQuery.ajax({
            url: API_URL,
            type: 'POST',
            crossDomain: true,
            data: JSON.stringify({Category: "furniture", Item: "example lamp", ItemDesc: "use", Seller: "Mr. Lamp", Email: "newemail@newemail.com", Price: "$5", Model: "lamp_mod_username", Texture: "lamp_tex_username"}),//apiData,
            dataType: 'json',
            contentType: "application/json",
            success: function (response) {
                //var resp = JSON.parse(response)
                alert(resp.status);
            },
            error: function (xhr, status) {
                alert("error");
            }
        });

        return false;
    });*/
        //Assign the username to the suitable fields and make sure if something changes to update that
        jQuery('#modUsername').val(jQuery('input[name=seller]').val());
        jQuery('#texUsername').val(jQuery('input[name=seller]').val());
        jQuery('input[name=seller]').on("change", updateUsername);

        //This is triggered when we click on the submit button
        jQuery("#sendAjax").click(function(){
            var form = document.forms['submitForm'];
            //e.preventDefault();

            //Please put here your API URL
            var API_URL = 'http://40.121.206.106:5073/additem';

            //Retrieve the model's name and put it in the model input
            jQuery('input[name=model]').val(jQuery("#imageName").val());
            //Update the texture filename
            jQuery('input[name=texture]').val(jQuery("#textureName").val());

            //collect the form data
            var data = {};
            for (var i = 0, ii = form.length; i <ii; ++i) {
                var input = form[i];
                if (input.name) {
                    data[input.name] = input.value;
                }
            }

            //var serialized = jQuery('form[name=submitForm]').serializeFormJSON();

            jQuery.ajax({
                url: API_URL,
                type: 'POST',
                crossDomain: true,
                data: JSON.stringify(data),
                dataType: 'text',
                contentType: "application/json",
                success: function (response) {
                    //var resp = JSON.parse(response);
                    //console.log(resp.statusText);
                    alert('Post Successful');
                },
                error: function (/*xhr,*/ status) {
                    //console.log(status);

                    alert("Post failed for some reason");
                }
            });

        });

        //File Upload for the model
        jQuery('.image-button').click(function(){
            jQuery('#fileupload').click();
        });

        //File Upload for the texture
        jQuery('.image-buttonx').click(function(){
            jQuery('#textureupload').click();
        });

        
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
                    image_src   = 'http://www.nebulashop.net/uploads/'+file.name;
                    original    = image_src.replace('_mod_'+jQuery('input[name=seller]').val(), '');

                    jQuery("#imageName").val(image_src);
                    jQuery("#modelName").html(original);
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

        //'use strict';
        // Change this to the location of your server-side upload handler:
        var url = './upload_texture.php';

        jQuery('#textureupload').fileupload({
            url: url,
            dataType: 'json',
            maxNumberOfFiles: 1,
            done: function (e, data) {
                jQuery.each(data.result.files, function (index, file) {

                    //Save the uploaded image file name to be used for later
                    image_src = 'http://www.nebulashop.net/uploads/'+file.name;
                    original    = image_src.replace('_tex_'+jQuery('input[name=seller]').val(), '');

                    jQuery("#textureName").val(image_src);
                    jQuery("#texName").html(original);
                });
            },
            progressall: function (e, data) {
                var progress = parseInt(data.loaded / data.total * 100, 10);
                jQuery('#progressx .progress-bar').css(
                    'width',
                    progress + '%'
                );
            }
        }).prop('disabled', !jQuery.support.fileInput)
            .parent().addClass(jQuery.support.fileInput ? undefined : 'disabled');
    });
