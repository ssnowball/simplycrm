<template>
  <div class="container new-width">
    <h3 v-if="!updateID">Create {{table}}</h3>
    <h3 v-if="updateID">Update {{table}}</h3>
    <form>
      <div class="row">
              <div class="input-field col s4">
                <input id="fName" type="text" v-model.trim="$v.name.$model">
                <label for="fName" v-bind:class="{ active: name }">Name</label>
                <div class="helper-text" v-if="$v.name.$invalid">*required</div>
              </div>
              <div class="input-field col s4">
                <input id="fDescription" type="text" v-model.trim="$v.description.$model">
                <label for="fDescription" v-bind:class="{ active: description }">Description</label>
                <div class="helper-text" v-if="$v.description.$invalid">*required</div>
              </div>
              <div class="input-field col s4">
                <input id="fCode" type="text" v-model.trim="$v.code.$model">
                <label for="fCode" v-bind:class="{ active: code }">Code (if have one)</label>
                <div class="helper-text" v-if="$v.code.$invalid">*required</div>
                
              </div>
      </div>
      <div class="row">
        <div class="input-field col s4">
          <input id="fCost" type="number" min=0  step="0.01" lang="nb" v-model.trim="$v.cost.$model">
          <label for="fCost" v-bind:class="{ active: cost }">Cost</label>
          <div class="helper-text" v-if="$v.cost.$invalid">*required</div>
        </div>
        
        <div class="input-field col s4">
          <input id="fCategory" type="text" v-model.trim="$v.category.$model">
          <label for="fCategory" v-bind:class="{ active: category }">Category</label>
          <div class="helper-text" v-if="$v.category.$invalid">*required</div>
        </div>
        <div class="input-field col s4">
          <input id="fSubCategory" type="text" v-model.trim="$v.subCategory.$model">
          <label for="fSubCategory" v-bind:class="{ active: subCategory }">Sub Category</label>
          <div class="helper-text" v-if="$v.subCategory.$invalid">*required</div>
        </div>

      </div>

      <div class="row">
        <div class="input-field col s4">
          <input type="hidden" id="fDocID" disabled v-model.trim="$v.docID.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fCreatedDate" disabled v-model.trim="$v.createdDate.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fCreatedUsername" disabled v-model.trim="$v.createdUsername.$model">
        </div>
      </div>
        
        <a class="waves-effect waves-light btn entryButton" v-on:click="changePage">cancel</a>
        <a v-if="!updateID" class="waves-effect waves-light btn entryButton" id="btn-create" v-bind:class="{ disabled: $v.GroupValid.$invalid }" v-on:click="createData('btn-create')">create</a>
        <a v-if="updateID" data-target="modal-delete" class="btn modal-trigger delete-color entryButton" id="delete-trigger">delete</a>
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
import {required, minValue} from 'vuelidate/lib/validators';

export default {
  name: 'productEntry',
  props: ["table", "updateID"],
  data() {
    return {
      name: "",
      description: "",
      code: "",
      cost: "",
      category: "",
      subCategory: "",
      docID: "",
      createdDate: "",
      createdUsername: "",
    };
  },
  validations: {
    name: {
      required,
    },
    cost: {
      required,
      minValue: minValue(0.01)
    },
    category: {
      required,
    },
    subCategory: {
      required,
    },
    description: {
      // required,
    },
    code: {
      // required,
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
    GroupValid: ['name', 'cost', 'category', 'subCategory'],
  },
  mounted: function () {
    M.AutoInit();
    if (this.updateID) {
      this.getUpdateData(this.updateID);
    }    
},
  methods: {
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
          
        });

    },
    getUpdateData(docID) {

      window.backend.getUpdateData(this.table, docID).then(result => {       

        if (result.Data === null) {
            M.toast({html: result.ReturnErr, classes: 'rounded'});
          } else {

            this.name = result.Data.Name;
            this.description = result.Data.Description;
            this.category = result.Data.Category;
            this.subCategory = result.Data.SubCategory;
            this.cost = result.Data.Cost;
            this.code = result.Data.Code;
            this.docID = result.Data.DocID;
            this.createdDate = result.Data.CreatedDate;
            this.createdUsername = result.Data.CreatedUsername;

            M.updateTextFields();
            // M.AutoInit();
          }
      });
    },
    createData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Description = self.description
      args.Category = self.category;
      args.SubCategory = self.subCategory;
      args.Cost = self.cost;
      args.Code = self.code;  

      window.backend.createData(self.table, args).then(result => {
        if (result.Success === true) {
          M.toast({html: "Added Item", classes: 'rounded'});
          this.changePage();
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
        document.getElementById("loader-show").classList.toggle('hide');

      });

    },
    updateData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Description = self.description
      args.Category = self.category;
      args.SubCategory = self.subCategory;
      args.Cost = self.cost;
      args.Code = self.code;
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

      });

    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
