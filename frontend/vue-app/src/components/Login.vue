<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-5">Login</h1>
                    <hr>
                        <form-tag @submitLoginEvent="submitHandler" name="myform" event="submitLoginEvent">
                        <text-input v-model="email" label="Email" type="email" name="email" required="true" ></text-input>
                        <text-input v-model="password" label="Password" type="password" name="password" required="true" ></text-input>

                    <hr>
                <input type="submit" class="btn btn-primary" value="Login">
                </form-tag>
            </div>
        </div>
    </div>
</template>


<script>
import TextInput from './forms/TextInput.vue'
import FormTag from './forms/FormTag.vue'
import { store } from './store.js'
import router from './../router/index.js'
import notie from 'notie'
import Security from './security.js'



export default {
    name: 'login',
    data(){
        return {
            email: '',
            password: '',
            store,
        }
    },
    components: {
        TextInput,
        FormTag
    },
    methods: {
        submitHandler(){
            const payload = {
                email: this.email,
                password: this.password
            }
            
            fetch(process.env.VUE_APP_API_URL + "/users/login", Security.requestOptions(payload))
            .then(res => res.json())
            .then(res => {
                if(res.error) {
                    notie.alert({
                        type: 'error',
                        text: res.message,
                        // stay: true,
                        // position: 'bottom'
                    })
                }else{
                    console.log(res)
                    store.token = res.data.token.token
                    store.user = {
                        id: res.data.user.id,
                        first_name: res.data.user.first_name,
                        last_name: res.data.user.last_name,
                        email: res.data.user.email
                    }
                    // save info to cookies
                    let date = new Date();
                    let expDays = 1;
                    date.setTime(date.getTime() + (expDays * 24 * 60 * 60 * 1000))
                    const expires = "expires=" + date.toUTCString()
                    // set the cookie
                    document.cookie = "_site_data="
                    + JSON.stringify(res.data)
                    + "; "
                    + expires
                    + "; path=/; SameSite=strict; Secure;"

                    router.push("/")
                }
            })
        }
    }
}
</script>