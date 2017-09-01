define(['vue'], function(Vue) {

    return Vue.extend({
        template:templateLoader("pub_sign_in.htm"),

        data: function() {
            return {
                signInModel:{},
            }
        },
        methods:{
            submitForm(e) {
                this.$validator.validateAll();
                if (!this.errors.any()) {
                    var self = this;
                    customerService.postUnsecured('/api/sign-in', this.signInModel,
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