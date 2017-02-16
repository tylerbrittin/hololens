
 
 

//check if a Var exists
function isDefined(x) {
    var undefined;
    return x !== undefined;
}



//New definitions
(function($) {
    $.fn.ajaxify = function(options) {
        var defaults = {
            /*sys_message: '#sys_message',
             message_alert: '#message-alert',*/
            no_html: 'yes'
        };
        this.each(function() {
            var $this = $(this);
            $this.submit(function()
            {
                var postdata = $this.serialize() + '&no_html=yes';
                var action = $this.attr("action");

                $.post(action, postdata,
                    function(data) {
                        //alert("Data Loaded: " + data);
                        var params  = data.split("||"),
                            message     = null,
                            silent_mode = false;

                        //Sometimes, we don't need to return a message, so we check if there exists a noreturn class
                        if (!$this.hasClass('noreturn'))
                        {
                            //Check whether it's JSON or not
                            try
                            {
                                var json    = $.parseJSON(data);

                                //Silent mode is used to return values but not to display them. The response can be used in the debugging
                                if(json && isDefined(json.silent_mode) && json.silent_mode == 'true') {
                                    silent_mode = true;
                                }

                                if(silent_mode == false)
                                {
                                    $.jGrowl(json.msg);
                                }

                                //In case we've got a returned callback JS function to be executed, then we'll do that.
                                if (json && isDefined(json.callback))
                                {
                                    eval(json.callback);
                                }
                            }
                            catch(e)
                            {
                                //console.log(e);
                                console.log(data);
                            }
                        };
                    });

                return false;
            });
        });
    };

    $.fn.extend({
        remove_text: function(str) {
            return this.each(function(){
                var self	= $(this);

                var reg		= str;
                if(self.html().indexOf(str) != -1) var txt	= self.html().replace(str,"");
                else var txt		= self.html();

                self.html(txt);
            });
        }
    });

    $.fn.extend({
        add_text: function(str) {
            return this.each(function(){
                var self	= $(this);
                var txt		= str+self.html();
                self.html(txt);
            });
        }
    });

    //Extend the Ajaxify Function to refresh the page after sending the request
    $.fn.ajaxify_refresh = function()
    {
        this.each(function()
        {
            var id = '#'+jQuery(this).attr('id');

            $(id).submit(function()
            {
                var postdata = $(id).serialize() + '&no_html=yes';
                var action = $(id).attr("action");

                $.post(action, postdata,
                    function(data)
                    {
                        document.location.reload();
                    });

                return false;
            });
        });
    }
})(jQuery);

$(function() {
    jQuery(".ajaxify").ajaxify();
    jQuery(".ajaxify_refresh").ajaxify_refresh(); //Used for functions that require refreshing the page like the Logout function
});