                                var form = document.forms['form'];

                                form.onsubmit = function (e) {
                                //stop regular form submission
                                e.preventDefault();

                                //collect the form data 
                                var data = {};
                                for (var i = 0, ii = form.length; i <ii; ++i) {
                                var input = form[i];
                                if (input.name) {
                                data[input.name] = input.value;
                            }
                        }

           //console.log(JSON.stringify(data));
                        //construct an HTTP request
                        var url = 'http://40.87.66.169:5073/additem'

                            $.ajax({
                                url: url,
                                type: 'POST',
                                crossDomain: true,
                                data: JSON.stringify(data),
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
                        //console.log(JSON.stringify(data))
                        //console.log(url);

                    (function (window) {
  function CorsAjax() {
    this.post = function(url, data, callback) {
      $.support.cors = true;
      var jqxhr = $.post(url, data, callback, "json")
             .error(function(jqXhHR, status, errorThrown) {
               if ($.browser.msie && window.XDomainRequest) {
                   var xdr = new XDomainRequest();
                   xdr.open("post", url);
                   xdr.onload = function () {
                    if (callback) {
                     callback(
                      JSON.parse(this.responseText), 
                      'success');
                    }
                  };
                  xdr.send(data);
               } else {
                 alert("corsAjax.post error: " + status + ", " + errorThrown);
               }
      });
    };

    this.get = function(url, callback) {
        $.support.cors = true;
        var jqxhr = $.get(url, null, callback, "json")
               .error(function(jqXhHR, status, errorThrown) {
                  if ($.browser.msie && window.XDomainRequest) {
                    var xdr = new XDomainRequest();
                    xdr.open("get", url);
                    xdr.onload = function () {
                     if (callback) {
                      callback(
                       JSON.parse(this.responseText), 
                       'success');
                     }
                    };
                    xdr.send();
                  } else {
                    alert("corsAjax.get error: " + status + ", " + errorThrown);
                  }
                });
        };
    };

    window.corsAjax = new CorsAjax();
})(window);

                };

