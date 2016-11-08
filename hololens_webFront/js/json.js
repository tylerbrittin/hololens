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

                        //construct an HTTP request
                        var xhr = new XMLHttpRequest();
                        xhr.open(form.method, form.action, true);
                        xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');

                        //send the collected data as JSON
                        xhr.send(JSON.stringify(data));

                        console.log(JSON.stringify(data));

                        xhr.onloadend = function () {
                        //done
                    };
                };