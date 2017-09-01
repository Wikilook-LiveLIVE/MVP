define([
    '/ui/htm/pub_sign_up.js',

], function(pubSignUp){
    return [
        {path:'/', component: pubSignUp, meta: {withToken: false}},
        {path:'/sign-up', component: pubSignUp, meta: {withToken: false}}
    ]
});