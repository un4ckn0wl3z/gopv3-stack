<template>
        <div class="form-check">
            <input class="form-check-input" :checked="checked" :required="required" :value="value" :name="name" :id="name" type="checkbox">
            <label class="form-check-label" :for="name">{{label}}</label>
        </div>
</template>


<script>
export default {
    name: 'CheckInput',
        props: [
        "value", "name", "required", "label", "checked"
    ],

}
</script>