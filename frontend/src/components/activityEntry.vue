<template>
  <div class="container new-width">
    <h3 v-if="!updateID">Create {{table}}</h3>
    <h3 v-if="updateID">Update {{table}}</h3>
    <form>
      <div class="row">
        <div class="input-field col s6">
        <select id="fType" v-model.trim="$v.rType.$model">
          <option value="" disabled selected>Choose your option</option>
          <option v-for="item in dropType" 
                  v-bind:key="item.key" 
                  v-bind:value="item.value">
          {{ item.value }}
          </option>
        </select>
        <label>Type</label>
        <div class="helper-text" v-if="$v.rType.$invalid">*required</div>
      </div>

      <div class="input-field col s5">
        <input id="fReferenceName" type="text" v-model.trim="$v.referenceName.$model">
        <input id="fReferenceID" type="hidden" disabled v-model.trim="$v.referenceID.$model">
        <label for="fReferenceName" v-bind:class="{ active: referenceName }">Reference</label>
        <div class="helper-text" v-if="$v.referenceName.$invalid">*required</div>
      </div>
      <div class="input-field col s1">
        <a class="waves-effect waves-light btn-small" id="account-find" v-on:click="getSubItem()" v-bind:class="{ disabled: $v.rType.$invalid }">find</a>
      </div>
    </div>
<div class="row">
  <div class="input-field col s12">
    <select id="fActivity" v-model.trim="$v.activity.$model">
          <option value="" disabled selected>Choose your option</option>
          <option v-for="item in dropActivity" 
                  v-bind:key="item.key" 
                  v-bind:value="item.value">
          {{ item.value }}
          </option>
        </select>
        <label>Activity</label>
        <div class="helper-text" v-if="$v.activity.$invalid">*required</div>
  </div>
</div>
<div class="row">
  <div class="input-field col s12">
    Detail
    <textarea id="fDetail" class="materialize-textarea" v-model.trim="$v.detail.$model"></textarea>
    <div class="helper-text" v-if="$v.detail.$invalid">*required</div>
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

      <!-- Modal Structure -->
      <div id="modal-subView" class="modal bottom-sheet">
        <div class="modal-content">
          <a data-target="modal-subView" class="btn modal-close">close</a> 
          <a class="waves-effect waves-light btn-small" v-on:click="removeSubItem()" v-bind:class="{ disabled: !referenceName }">Clear Selected</a> 
          <a class="waves-effect waves-light btn-small" v-on:click="filterResultsAccount()" v-bind:class="{ disabled: !referenceName }">View Selected</a>
          <div style="margin-bottom: 10px">
          <el-row>
            <el-col :span="6">
              <el-input v-model="filters[0].value" placeholder="Search"></el-input>
            </el-col>
          </el-row>
        </div>
          <data-tables :data="subData" :pagination-props="{ pageSizes: [5, 10, 15] }" :filters="filters" :action-col="actionCol">
            <el-table-column v-for="title in titles" :prop="title.prop" :label="title.label" :key="title.label" sortable="custom">
            </el-table-column>
          </data-tables>

        </div>
      </div>

  </div>
</template>

<script>
import M from 'materialize-css'
import {required} from 'vuelidate/lib/validators';

export default {
  name: 'activityEntry',
  props: ["table", "updateID"],
  data() {
    return {
      dropType: [{ key: 'Account', value: 'Account' }, { key: 'Lead', value: 'Lead' }],
      dropActivity: [{ key: 'Note', value: 'Note' }, 
                      { key: 'Call', value: 'Call' }, 
                      { key: 'Email', value: 'Email' }, 
                      { key: 'Meeting', value: 'Meeting' }, 
                      { key: 'Quote', value: 'Quote' }, 
                      { key: 'Invoice', value: 'Invoice' }],
      rType: "",
      referenceName: "",
      referenceID: "",
      activity: "",
      detail: "",
      titles: [],
      filters: [
        {
          prop: '',
          value: ''
        }
      ],
      actionCol: {
        label: 'Actions',
        props: {
          align: 'center',
        },
        buttons: [{
          props: {
            type: 'primary',
          },
          handler: row => {
            // console.log(row);
            this.selectedSubItem(row);
          },
          label: 'Select'
        }]
      },
      subData: [],
      accountSearch: "",
    };
  },
  validations: {
    rType: {
      required,
    },
    referenceName: {
      required,
    },
    referenceID: {
      required,
    },
    activity: {
      required,
    },
    detail: {
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
    GroupValid: ['rType', 'referenceName', 'referenceID', 'activity', 'detail'],
  },
  mounted: function () {
    M.AutoInit();
    
    if (this.updateID) {
      this.getUpdateData(this.updateID);
    }    
},
watch: {
  rType: function () {
    if (this.updateID) {
      var s = document.getElementById('fType');
      var v = this.rType;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
    }
    },
    activity: function () {
      if (this.updateID) {
      var s = document.getElementById('fActivity');
      var v = this.activity;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
},
  methods: {
    filterResultsAccount() {
      this.$emit('filter', this.rType, 'list', this.referenceName);
    },
    changePage () {
      this.$emit('clicked', 'list')
    },
    getSubItem() {
      var self = this;

      // console.log(this.rType);
      window.backend.readSubData(self.rType).then(result => { 
        
        var returned = JSON.parse(JSON.stringify(result));

        if (returned.Data === null) {
            self.titles = returned.Titles;
            self.subData = [];
            M.toast({html: result.ReturnErr, classes: 'rounded'});

          } else {

            self.titles = returned.Titles;
            self.subData = returned.Data;
            self.accountSearch = self.rType;

            var filt = [];
            var i;
            for (i = 0; i < returned.Titles.length; i++) {
              filt.push(returned.Titles[i].prop);
            }
            self.filters[0].prop = filt;
          
            M.Modal.getInstance(document.getElementById('modal-subView')).open();
          
          }
      });
    },
    selectedSubItem(row) {
      M.Modal.getInstance(document.getElementById('modal-subView')).close();

      this.referenceID = row.DocID;
      this.referenceName = row.Name;

      M.updateTextFields();

    },
    removeSubItem() {
      this.referenceID = "";
      this.referenceName = "";

      M.Modal.getInstance(document.getElementById('modal-account')).close();
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

            this.rType = result.Data.Type;
            this.referenceName = result.Data.ReferenceName
            this.referenceID = result.Data.ReferenceID;
            this.activity = result.Data.Activity;
            this.detail = result.Data.Detail;

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
      args.Type = self.rType;
      args.ReferenceName = self.referenceName
      args.ReferenceID = self.referenceID;
      args.Activity = self.activity;
      args.Detail = self.detail;

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
      args.Type = self.rType;
      args.ReferenceName = self.referenceName
      args.ReferenceID = self.referenceID;
      args.Activity = self.activity;
      args.Detail = self.detail;
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
