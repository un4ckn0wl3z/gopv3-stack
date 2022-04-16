<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">All Users</h1>
            </div>
            <hr>
            <table v-if="ready" class="table table-compact table-striped">
                <thead>
                    <tr>
                        <th>User</th>
                        <th>Email</th>
                        <th>Active</th>
                        <th>Status</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="u in users" v-bind:key="u.id">
                        <td>
                            <router-link :to="`/admin/users/${u.id}`">{{u.first_name}}, {{u.last_name}}</router-link>
                        </td>
                         <td>
                            {{u.email}}
                        </td>
                         <td v-if="u.active === 1">
                            <span class="badge bg-success">Active</span>                            
                        </td>
                         <td v-else>
                            <span class="badge bg-danger">Inactive</span>                                                         
                        </td>                                                   
                        <td v-if="u.token.id > 0">
                            <a href="javascript:void(0);">
                                <span class="badge bg-success" @click="logUserOut(u.id)">Logged in</span>
                            </a>
                        </td>
                        <td v-else>
                            <span class="badge bg-danger" >Not logged in</span>
                        </td>
                    </tr>
                </tbody>
            </table>
            <p v-else >Loading...</p>
        </div>
    </div>
</template>

<script>

import Security from './security.js'
import {store} from './store.js'
import notie from 'notie'



export default {
    data(){
        return {
            users: [],
            ready: false
        }
    },
    beforeMount(){
        Security.requireToken()
        
        fetch(process.env.VUE_APP_API_URL + "/admin/users", Security.requestOptions({}))
        .then(res => res.json())
        .then(res => {
            if(res.error){
                this.$emit('error', res.message)
            }else {
                this.users = res.data.users
                this.ready = true
            }
        })
        .catch(err => {
            this.$emit('error', err)
        })
    },
    methods: {
        logUserOut(id){
            if (id !== store.user.id){
                notie.confirm({
                    text: 'Are you sure you want to log this user out?',
                    submitText: 'Log Out',
                    submitCallback: () => {
                        fetch(`${process.env.VUE_APP_API_URL}/admin/log-user-out/${id}`, Security.requestOptions({}))
                        .then(res => res.json())
                        .then(res => {
                            if(res.error){
                                this.$emit('error', res.message)
                            }else {
                                this.$emit('success', 'logged user out!')   
                                this.$emit('forceUpdate')                                         
                            }
                        })
                        .catch(err => {
                                this.$emit('error', err)            
                        })                         
                    }
                })
            }else {
                this.$emit('error', "You can't log yourself out!")
            }
        }
    }
}
</script>