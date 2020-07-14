<template>
  <div class="container new-width">
    <h3>{{table}}</h3>

    <a v-if="ShowCreateButton" href="#" class="waves-effect waves-light btn" v-on:click="changePage">create</a>

    <div style="margin-bottom: 10px">
      <el-row>
        <el-col :span="6">
          <el-input v-model="filters[0].value" placeholder="Search"></el-input>
        </el-col>
      </el-row>
    </div>

    <perfect-scrollbar>
        <data-tables :data="data" :pagination-props="{ pageSizes: [5, 10, 15] }" :filters="filters" :action-col="actionCol" layout="pagination, table">
      <el-table-column v-for="title in titles" :prop="title.prop" :label="title.label" :key="title.label" sortable="custom">
      </el-table-column>
    </data-tables>
    </perfect-scrollbar>
    

   <div class="progress hide" id="loader-show">
        <div class="indeterminate"></div>
      </div>

  </div>
</template>

<script>

import M from 'materialize-css'


export default {
  name: 'list',
  props: ["table", "filteredValue"],
  data() {
    return {
      ShowCreateButton: false,
      data: [],
      titles: [],
      layout: 'pagination, table',
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
            // icon: 'el-icon-edit'
          },
          handler: row => {
            if (this.table != "Audit") {
              this.UpdateData(this.data[this.data.indexOf(row)].DocID);
            }
          },
          label: 'Expand'
        }]
      },
    };
  },
  mounted: function () {
    M.AutoInit();
    this.readData();
    if (this.filteredValue) {
      this.filterData();
    }
    

    // if (this.table === "Audit" || this.table === "Option") {
    //   this.ShowCreateButton = false;
    // }
},
watch: {
    // whenever question changes, this function will run
    table: function () {
      this.readData();
    },
    filteredValue: function() {
      this.filters[0].value = this.filteredValue;
    },
},
  methods: {
    filterData() {
      this.filters[0].value = this.filteredValue;
    },
    changePage () {
      this.$emit('clicked', this.table + 'Entry')
    },
    UpdateData (docID) {
      this.$emit('UpdateData', this.table + 'Entry', docID)
    },
    readData: function() {
      var self = this;

        window.backend.readData(self.table).then(result => {
          var returned = JSON.parse(JSON.stringify(result));

          switch(self.table) {
            case "Option":
              if (returned.Data === null) {
                self.ShowCreateButton = true;
              } else {
                self.ShowCreateButton = false;
              }
              
              break;
            case "Audit":
              self.ShowCreateButton = false;
              break;
            default:
              self.ShowCreateButton = true;
          }

          if (returned.Data === null) {
            self.titles = returned.Titles;
            self.data = [];

          } else {

            self.titles = returned.Titles;
            self.data = returned.Data;

            var filt = [];
            var i;
            for (i = 0; i < returned.Titles.length; i++) {
              filt.push(returned.Titles[i].prop);
            }
            self.filters[0].prop = filt;
          }
          
          
        });
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.ps {
  height: 500px;
}
</style>