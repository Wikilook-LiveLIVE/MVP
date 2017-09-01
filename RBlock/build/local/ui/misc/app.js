/*
    Utils
*/
function templateLoader (tempName){
    return $.ajax({
        type: "GET",
        url: "/ui/htm/"+tempName,
        cache: false,
        async: false
    }).responseText;
}


var customerService = {
    //global app vue object
    appVue: null,

    init: function (vue){
        this.appVue = vue;
    },


    postUnsecured: function(url, model, successFn, errorFn){
        $.ajax({
            type: "POST", url: url,
            dataType: "json", contentType: "application/json;charset=utf-8",
            cache: false,
            data: JSON.stringify(model),
            timeout: 5000,
            success: successFn, error: errorFn
        });
    },

    postSecured: function(url, model, successFn, errorFn){
        $.ajax({
            type: "POST", url: url,
            headers: {"token":this.appVue.user.token},
            dataType: "json", contentType: "application/json;charset=utf-8",
            cache: false,
            data: JSON.stringify(model),
            timeout: 5000,
            success: successFn, error: errorFn
        });
    },

    getSecured: function(url, successFn, errorFn){
        $.ajax({
            type: "GET", url: url,
            headers: {"token":this.appVue.user.token},
            dataType: "json", contentType: "application/json;charset=utf-8",
            cache: false,
            timeout: 5000,
            success: successFn, error: errorFn
        });
    }
}


var i18nService = {
    //global app vue object
    appVue: null,

    init: function (vue){
        this.appVue = vue;
    },
    //load language bundle
    getBundle: function (language) {
        var languageFilePath = '/ui/i18n/' + language + '.json?1.0';
        this.appVue.i18n = $.parseJSON($.ajax({
            type: 'GET',
            url: '/ui/i18n/' + language + '.json?1.0',
            cache: false,
            async: false,
            dataType: "json"
        }).responseText);
    },

    //add error to serverError
    convertError: function (response) {
        if (response.err != null) {
            console.log("i18nService convertError: "+response.err.code);
            this.appVue.serverError = this.appVue.i18n.error[response.err.code];
        }
    },

    plainError: function (key) {
        console.log("i18nService plainError: "+key);
        this.appVue.serverError = this.appVue.i18n.error[key];
    },
    //get from vocabulary
    get: function (key) {
        console.log("i18nService get: "+key);
        return this.appVue.i18n.vocabulary[key];
    }
}


require.config({
    baseUrl: "",
    paths: {
        'vue': '/ui/misc/lib/vue/vue.min',
        'vue_router': '/ui/misc/lib/vue/vue-router.min',
        'vee_validate': '/ui/misc/lib/vue/vee-validate.min'
    },
    shim: {
        vue: {exports: 'Vue'},
        vue_router: {exports: 'VueRouter'},
        vee_validate: {exports: 'VeeValidate'}
    }
});

require([
        'vue',
        'vue_router',
        'vee_validate',
        '/ui/misc/router.js'
    ], function(Vue, VueRouter, VeeValidate, AppRoutes){
        Vue.use(VeeValidate);
        Vue.use(VueRouter);

        //server error

        Vue.component('serverErr', {
            template: '<div class="errorFormContainer brand-danger" @dblclick="closeMsg" ' +
            ' v-show="$root.serverError != null" title="Double click to hide">' +
            '{{$root.serverError}}</div>',
            methods:{
                closeMsg(){
                    this.$root.serverError = null;
                }
            }

        });

        var router = new VueRouter({
            mode: 'hash',
            routes: AppRoutes
        });

        var vm = new Vue({
            el: '#app',

            data: function() {
                return {
                    serverError: null,
                    user: {
                        token: null,
                        lang: 'en'
                    },
                    i18n: null
                }
            },

            created: function(){
                customerService.init(this);
                customerService.tryRestoreSession();
                i18nService.init(this);
                i18nService.getBundle(this.user.lang);
            },

            mounted: function(){
                if (this.user.token == null) {
                    router.push('sign-up');
                }
            },

            methods: {
                signOut(){
                    customerService.destroySession();
                    router.push('sign-in');
                }
            },

            router: router
        });

        router.beforeEach(function(to, from, next) {
            if (to.meta.withToken){
                //valida token ?
                if (vm.user.token == null){
                    next(false);
                }else{
                    next(true);
                }
            }else{
                next(true);
            }
        });

    }
);

