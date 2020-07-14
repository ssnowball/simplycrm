<template>
  <div id="app">

     <div class="navbar-fixed">
      <nav class="new-back-colour"> 
        <div class="nav-wrapper">
          <div class="nav-logo"></div>
          <a href="#" class="brand-logo center"><img src="./assets/images/logo.png" alt="logo" height="50"></a>
          <a href="#" data-target="slide-out" class="left sidenav-trigger base-text-colour">Menu</a>
        </div>
      </nav>
    </div>

  <ul id="slide-out" class="sidenav">
        <li class="base-colour"><div class="user-view">
          <span class="white-text name">{{ownerName}}</span>
          <span class="white-text name">{{companyName}}</span>
        </div></li>
        <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Lead', loadedPage='list', filteredValue=''">Lead</a></li>
        <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Account', loadedPage='list', filteredValue=''">Account</a></li>
        <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Activity', loadedPage='list', filteredValue=''">Activity</a></li>

        <li><div class="divider divLine"></div></li>
          <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Product', loadedPage='list', filteredValue=''">Product</a></li>
          <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Campaign', loadedPage='list', filteredValue=''">Campaign</a></li>

          <li><div class="divider divLine"></div></li>
          <li><a class="waves-effect" href="#" v-on:click="loadedItem = 'Audit', loadedPage='list', filteredValue=''">Audit</a></li>
          <li><a class="waves-effect" href="#/" v-on:click="loadedItem = 'Option', loadedPage='list', filteredValue=''">Options</a></li>

      </ul>
      
      <productEntry v-if="loadedPage === 'ProductEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" />
      <campaignEntry v-if="loadedPage === 'CampaignEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" />
      
      <accountEntry v-if="loadedPage === 'AccountEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" @filter="onClickFilter" />
      <leadEntry v-if="loadedPage === 'LeadEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" @filter="onClickFilter" />
      <activityEntry v-if="loadedPage === 'ActivityEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" @filter="onClickFilter" />
      
      <optionEntry v-if="loadedPage === 'OptionEntry'" :table=loadedItem :updateID=updateID @clicked="onClickChild" @UpdateUser="onUserUpdate" />

      <list v-if="loadedPage === 'list'" :table=loadedItem :filteredValue=filteredValue @clicked="onClickChild" @UpdateData="onClickUpdate" />

      <setupLoad v-if="loadedPage === 'setup'" />

  </div>
</template>

<script>

import productEntry from "./components/productEntry.vue";
import campaignEntry from "./components/campaignEntry.vue";
import accountEntry from "./components/accountEntry.vue";
import leadEntry from "./components/leadEntry.vue";
import activityEntry from "./components/activityEntry.vue";
import optionEntry from "./components/optionEntry.vue";
import list from "./components/list.vue";
import setupLoad from "./components/setupLoad.vue";

import "./assets/css/main.css";
import M from 'materialize-css';

export default {
  name: "app",
  components: {
    productEntry,
    campaignEntry,
    accountEntry,
    leadEntry,
    activityEntry,
    optionEntry,
    list,
    setupLoad,
  },
  data() {
    return {
      loadedPage: "setup", // list | entries
      loadedItem: "Lead",
      updateID: "",
      filteredValue: "",
      ownerName: "",
      companyName: "",
    }
  },
  mounted: function () {
    M.AutoInit();
    this.loadDatase();
    var elems = document.querySelectorAll('.sidenav');
    // var options = {format:'dd/mm/yyyy'};
    M.Sidenav.init(elems);
},
watch: {
    // whenever question changes, this function will run
    loadedPage: function () {
      if (this.loadedPage != "setup") {
        this.getUserdata();
      }
      
      M.Sidenav.getInstance(document.getElementById('slide-out')).close();
    },
    loadedItem: function() {
      M.Sidenav.getInstance(document.getElementById('slide-out')).close();
    },
},
methods: {
  loadDatase() {
    var self = this;
    window.backend.databaseBuild().then(result => {
          
        if (result.Success === true) {
          if (result.ReturnErr != "Database Exists") {
            self.loadedPage = "OptionEntry";
            self.loadedItem = "Option";
          } else {
            self.loadedPage = "list";
            self.loadedItem = "Lead";
          }
          
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
            
        });

  },
  getUserdata() {
    var self = this;

    window.backend.readData("Option").then(result => {
          var returned = JSON.parse(JSON.stringify(result));

          if (returned.Data === null) {
            self.ownerName = "User Name";
            self.companyName = "Company Name";
          } else {
            self.ownerName = returned.Data[0].Name;
            self.companyName = returned.Data[0].Company;
          }
            
        });
  },
  onClickChild (value) {
      this.updateID = "";
      this.loadedPage = value;
      this.filteredValue = "";
    },
  onClickUpdate (value, docID) {
      this.loadedPage = value;
      this.updateID = docID;
      this.filteredValue = "";
    },
  onUserUpdate () {
      this.getUserdata();
    },
  onClickFilter(tableName, value, filterVal) { 
      this.updateID = "";
      this.loadedPage = value;
      this.loadedItem = tableName;
      this.filteredValue = filterVal;
  },
  },
};
</script>
