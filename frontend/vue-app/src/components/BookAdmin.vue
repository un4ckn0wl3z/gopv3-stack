<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">Manage Books</h1>
                <hr>

                <table v-if="ready" class="table table-striped table-compact">
                    <thead>
                        <tr>
                            <th>Book</th>
                            <th>Author</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="b in books" :key="b.id">
                            <td>
                                <router-link :to="`/admin/books/${b.id}`">{{b.title}}</router-link>
                            </td>
                            <td>{{b.author.author_name}}</td>
                        </tr>
                    </tbody>

                </table>

            </div>
        </div>
    </div>
</template>


<script>

import Security from './security.js'

export default {
    name: 'BooksAdmin',
    data() {
        return {
            books: [],
            ready: false,
        }
    },
    mounted() {
        Security.requireToken()
        
        fetch(process.env.VUE_APP_API_URL + "/books")
        .then(res => res.json())
        .then(res => {
            if (res.error){
                this.$emit('error', res.message)
            }else{
                this.books = res.data.books
                this.ready = true
            }
        })
    }
}
</script>