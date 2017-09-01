define(['vue'], function(Vue) {

    return Vue.extend({
        template:templateLoader("pub_sign_up.htm"),

        data: function() {
            return {
                signUpModel:{
                    login:"",
                    password:"",
                    confirmPassword:""
                }
            }
        },
        methods:{
            submitForm(e) {
                this.$validator.validateAll();
                if (!this.errors.any()) {
                    var self = this;
                    customerService.postUnsecured('/api/sign-up', this.signUpModel,
                        function(data, textStatus, request) {
                            customerService.updateSession(request);
                            if (data.err == null) {
                                self.$router.push('policies');
                            }else{
                                i18nService.convertError(data);
                            }
                        },
                        function(respose) {
                            //500err
                            i18nService.plainError("ERR_METHOD_UNAVAILABLE");
                        }
                    );
                }
            }
        }
    });
});