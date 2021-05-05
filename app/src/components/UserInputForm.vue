<template>
<div>
    <form @submit="submit($event)">
        <div class="field">
            <label class="label">{{editAddText}} Login ID</label>
            <div class="control">
                <input class="input" v-model="user.loginId" type="text" placeholder="LoginId" required>
            </div>
        </div>
        <div class="field">
            <label class="label">{{editAddText}} Full Name</label>
            <div class="control">
                <input class="input" v-model="user.fullname" type="text" placeholder="Full Name" required>
            </div>
        </div>
        <div class="field-body">
            <div class="field is-narrow">
            <label class="label">{{editAddText}} Enabled</label>    
            <div class="control">
                <label class="radio">
                <input type="radio" v-model="user.enabled" value="true" name='enabled' required>
                Yes
                </label>
                <label class="radio">
                <input type="radio" v-model="user.enabled" value="false" name='enabled'>
                No
                </label>
            </div>
            </div>
        </div>

        <div class="field is-grouped" style="margin-top:20px">
            <div class="control">
                <button class="button is-link">{{editAddText}} User</button>
            </div>
        </div>
    </form>
</div>
</template>
<script>
export default {
  name: 'UserInputForm',
  props: {
    editId: String,
    user: Object
  },
   computed: {
            editAddText(){
            return this.editId ? "Edit":"Insert";
            }
        }
    
}
</script>