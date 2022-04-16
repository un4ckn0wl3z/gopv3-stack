<template>
    <div class="container">
        <div class="row">
            <div class="col-md-2">
                <img v-if="ready" class="img-fluid img-thumbnail" :src="`${imgPath}/covers/${book.slug}.jpg`" alt="cover">
            </div>
            <div class="col-md-10">
                <template v-if="ready" >
                    <h3 class="mt-3">{{book.title}}</h3>
                    <hr>
                    <p>
                        <strong>Author:</strong> {{book.author.author_name}} <br>
                        <strong>Published:</strong> {{book.publication_year}} <br>
                    </p>
                    <p>
                        {{book.description}}
                    </p>                    
                </template>
                <p v-else >Loading...</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            book: {},
            imgPath: process.env.VUE_APP_IMAGE_URL,
            ready: false

        }
    },
    beforeMount(){
        fetch(process.env.VUE_APP_API_URL + "/books/" + this.$route.params.bookName)
        .then(res => res.json())
        .then(res => {
          if (res.error) {
              this.$emit('error', res.message)
          } else {
              this.book = res.data
              this.ready = true
          }
        }).catch(err => {
            this.$emit('error', err)
        })
    }
}
</script>