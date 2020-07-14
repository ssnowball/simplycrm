<template>
  <div class="container new-width">
    <h3 v-if="!updateID">Create {{table}}</h3>
    <h3 v-if="updateID">Update {{table}}</h3>
    <h5>You can enter your company details here</h5>
    <form>
      <div class="row">
        <div class="input-field col s6">
          <input id="fName" type="text" v-model.trim="$v.name.$model">
          <label for="fName" v-bind:class="{ active: name }">Name</label>
          <div class="helper-text" v-if="$v.name.$invalid">*required</div>
        </div>
        <div class="input-field col s6">
          <input id="fCompnay" type="text" v-model.trim="$v.company.$model">
          <label for="fCompnay" v-bind:class="{ active: company }">Company</label>
          <div class="helper-text" v-if="$v.company.$invalid">*required</div>
        </div>
</div>
<div class="row">
        <div class="input-field col s4">
          <input type="hidden" id="fDocID" disabled v-model.trim="$v.docID.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fcreatedDate" disabled v-model.trim="$v.createdDate.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fcreatedUsername" disabled v-model.trim="$v.createdUsername.$model">
        </div>
      </div>

          <a class="waves-effect waves-light btn entryButton" v-on:click="changePage">cancel</a>
          <a v-if="updateID" data-target="modal-delete" class="btn modal-trigger delete-color entryButton" id="delete-trigger">delete</a>
          <a v-if="!updateID" class="waves-effect waves-light btn entryButton" id="btn-create" v-bind:class="{ disabled: $v.GroupValid.$invalid }" v-on:click="createData('btn-create')">create</a>
          <a v-if="updateID" class="waves-effect waves-light btn entryButton" id="btn-update" v-bind:class="{ disabled: $v.GroupValid.$invalid }" v-on:click="updateData('btn-update')">update</a>
        
        
        
        
        
        

      </form>

      <div class="progress hide" id="loader-show">
        <div class="indeterminate"></div>
      </div>

      <!-- Modal Structure -->
      <div id="modal-delete" class="modal">
        <div class="modal-content">
          <h4>Confirm Delete</h4>
          <p>You are about to delete please confirm</p>
        </div>
        <div class="modal-footer">
          <a href="#!" class="modal-close waves-effect waves-green btn-flat">Cancel</a>
          <a class="waves-effect waves-light btn" id="btn-delete" v-on:click="deleteData()">Confirm</a>
        </div>
      </div>

  </div>
</template>

<script>
import M from 'materialize-css'
import {required} from 'vuelidate/lib/validators';

export default {
  name: 'optionEntry',
  props: ["table", "updateID"],
  data() {
    return {
      name: "",
      company: "",
      createdDate: "",
      createdUsername: "",
    };
  },
  validations: {
    name: {
      required,
    },
    company: {
      required,
    },
    docID: {
      // required,
    },
    createdDate: {
      // required,
    },
    createdUsername: {
      // required,
    },
    GroupValid: ['name', 'company'],
  },
  mounted: function () {
    M.AutoInit();
    
    if (this.updateID) {
      this.getUpdateData(this.updateID);
    }    
},
  methods: {
    UpdateUser () {
      this.$emit('UpdateUser')
    },
    changePage () {
      this.$emit('clicked', 'list')
    },
    deleteData: function() {
      document.getElementById("loader-show").classList.toggle('hide');

      var self = this;

      window.backend.deleteData(self.table, self.updateID).then(result => {

        if (result.Success === true) {
          M.toast({html: "Deleted Item", classes: 'rounded'});   
          this.changePage(); 
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }

        document.getElementById("loader-show").classList.toggle('hide');
        self.UpdateUser();
          
        });

    },
    getUpdateData(docID) {

      window.backend.getUpdateData(this.table, docID).then(result => {       

        if (result.Data === null) {
            M.toast({html: result.ReturnErr, classes: 'rounded'});
          } else {

            this.name = result.Data.Name;
            this.company = result.Data.Company;
            this.docID = result.Data.DocID;
            this.createdDate = result.Data.CreatedDate;
            this.createdUsername = result.Data.CreatedUsername;

            M.updateTextFields();
          }
      });
    },
    createData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Company = self.company

      window.backend.createData(self.table, args).then(result => {
        if (result.Success === true) {
          M.toast({html: "Added Item", classes: 'rounded'});
          this.changePage();
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
        document.getElementById("loader-show").classList.toggle('hide');
        self.UpdateUser();

      });

    },
    updateData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Company = self.company
      args.DocID = self.docID;
      args.CreatedDate = self.createdDate;
      args.CreatedUsername = self.createdUsername;

      window.backend.updateData(self.table, args).then(result => {
        if (result.Success === true) {
          M.toast({html: "Updated Item", classes: 'rounded'}); 
          this.changePage(); 
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
        document.getElementById("loader-show").classList.toggle('hide');
        self.UpdateUser();

      });

    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
