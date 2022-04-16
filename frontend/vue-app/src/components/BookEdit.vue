<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">Add/Edit Book</h1>
                <hr>

                <form-tag @bookEditEvent="submitHandler" name="bookForm" event="bookEditEvent">

                    <div v-if="book.slug !== ''" class="mb-3">
                        <img class="img-fluid img-thumbnail book-cover" :src="`${imgPath}/covers/${book.slug}.jpg`" alt="cover">
                    </div>

                    <div class="mb-3">
                        <label for="formFile" class="form-label">Cover Image</label>
                        <input v-if="book.id === 0" ref="coverInput" class="form-control" type="file" id="formFile" 
                            required accept="image/jpg" @change="loadCoverImage">

                        <input v-else ref="coverInput" class="form-control" type="file" id="formFile" 
                            accept="image/jpg" @change="loadCoverImage">
                    </div>

                    <text-input 
                        v-model="book.title" 
                        type="text" 
                        required="true" 
                        label="Title" 
                        :value="book.title" 
                        name="title"
                    ></text-input>

                    <select-input name="author-id" v-model="book.author_id" :items="authors" required="required" label="Athor">

                    </select-input>

                    <text-input 
                        v-model="book.publication_year" 
                        type="number" 
                        required="true" 
                        label="Publication Year" 
                        :value="book.publication_year" 
                        name="publication-year"
                    ></text-input>

                    <div class="mb-3">
                        <label for="description" class="form-label">Description</label>
                        <textarea required v-model="book.description" class="form-control" id="description" rows="3"></textarea>
                    </div>


                    <div class="mb-3">
                        <label for="genres" class="form-label">Genres</label>
                        <select ref="genres" class="form-select" id="genres" required size="7" v-model="book.genre_ids" multiple>
                            <option v-for="g in genres" :value="g.value" :key="g.value">{{g.text}}</option>
                        </select>
                        
                    </div>   

                    <hr>

                    <div class="float-start">
                        <input type="submit" class="btn btn-primary me-2" value="Save" />
                        <router-link to="/admin/books" class="btn btn-outline-secondary">Cancel</router-link>
                    </div>
                    <div class="float-end">
                        <a v-if="book.id > 0" class="btn btn-danger" href="javascript:void(0);" @click="confirmDelete(book.id)">Delete</a>
                    </div>
                    <div class="clearfix"></div>

                </form-tag>

            </div>
        </div>
    </div>
</template>

<script>

import Security from './security.js'
import TextInput from './forms/TextInput.vue'
import FormTag from './forms/FormTag.vue'
import SelectInput from './forms/SelectInput.vue'
import router from './../router/index.js'

import notie from 'notie'


export default {
    beforeMount(){
        Security.requireToken()

        // get book for edit
        if (parseInt(String(this.$route.params.bookId), 10) > 0){
            // edit
            fetch(`${process.env.VUE_APP_API_URL}/admin/books/${this.$route.params.bookId}`, Security.requestOptions({}))
            .then(res => res.json())
            .then(res => {
                if(res.error){
                    this.$emit('error', res.message)
                }else {
                    this.book = res.data             
                }
            })
            .catch(err => {
                    this.$emit('error', err)            
            })              

   
        } 
            // add
            fetch(`${process.env.VUE_APP_API_URL}/admin/authors/all`, Security.requestOptions({}))
            .then(res => res.json())
            .then(res => {
                if(res.error){
                    this.$emit('error', res.message)
                }else {
                    this.authors = res.data             
                }
            })
            .catch(err => {
                    this.$emit('error', err)            
            })  

        

    },
    data() {
        return {
            book: {
                id: 0,
                title: '',
                author_id: 0,
                publication_year: null, 
                description: '',
                cover: '',
                slug: '',
                genres: [],
                genre_ids: []
            },
            authors: [],
            imgPath: process.env.VUE_APP_IMAGE_URL,
            genres: [
                {value: 1, text: 'Science Fiction'},
                {value: 2, text: 'Fantasy'},
                {value: 3, text: 'Romance'},
                {value: 4, text: 'Thriller'},
                {value: 5, text: 'Mystery'},
                {value: 6, text: 'Horror'},
                {value: 7, text: 'Classic'},
            ],
        }
    },
    components: {
        'form-tag': FormTag,
        'text-input': TextInput,
        'select-input': SelectInput,
    },
    methods: {
        submitHandler() {
            const payload = {
                id: this.book.id,
                title: this.book.title,
                author_id: parseInt(this.book.author_id, 10),
                publication_year: parseInt(this.book.publication_year, 10), 
                description: this.book.description,
                cover: this.book.cover,
                slug: this.book.slug,
                genre_ids: this.book.genre_ids,
            }
            
            fetch(process.env.VUE_APP_API_URL + "/admin/books/save", Security.requestOptions(payload))
            .then(res => res.json())
            .then(res => {
                if(res.error){
                    this.$emit('error', res.message)
                } else {
                    this.$emit('success', 'Changes saved')
                    router.push("/admin/books")
                }
            })
            .catch(err => {
                this.$emit('error', err)
            })


        },
        loadCoverImage() {
            // get image ref
            const file = this.$refs.coverInput.files[0];
            // encode the files
            const reader = new FileReader()
            reader.onloadend = () => {
                const base64String = reader.result.replace("data:", "").replace(/^.+,/, "")
                this.book.cover = base64String
                // alert(base64String)
            }

            reader.readAsDataURL(file)

        },
        confirmDelete(id) {
            notie.confirm({
                text: 'Are you sure you want to delete this book?',
                submitText: 'Delete',
                submitCallback: () => {
                    let payload = {
                        id: id
                    }

                    fetch(process.env.VUE_APP_API_URL + "/admin/books/delete", Security.requestOptions(payload))
                    .then(res => res.json())
                    .then(res => {
                        if(res.error){
                            this.$emit('error', res.message)
                        } else {
                            this.$emit('success', "Book deleted!")  
                            router.push("/admin/books")                     
                        }
                    })
                    .catch(err => {
                        this.$emit('error', err)               
                    })

                }
            })
        }
    },
    beforeUpdate() {
            this.book = {
                id: 0,
                title: '',
                author_id: 0,
                publication_year: 0, 
                description: '',
                cover: '',
                slug: '',
                genres: [],
                genre_ids: []
            }

            this.authors = [],
            this.imgPath = process.env.VUE_APP_IMAGE_URL
            this.genres = [
                {value: 1, text: 'Science Fiction'},
                {value: 2, text: 'Fantasy'},
                {value: 3, text: 'Romance'},
                {value: 4, text: 'Thriller'},
                {value: 5, text: 'Mystery'},
                {value: 6, text: 'Horror'},
                {value: 7, text: 'Classic'},
            ]
    }
}
</script>

<style scoped>
.book-cover {
    max-width: 10em;
}
</style>