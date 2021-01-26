<template>
    <div id="app">
        <div class="left" style="width: 20%;height: 100%;position: fixed">
            <el-scrollbar style="height: 100%">
                <el-tree :data="data" :props="defaultProps" lazy :load="handleNodeClick" @node-click="handleNodeClick1"></el-tree>
            </el-scrollbar>
        </div>
        <div class="right" style="margin-left: 20%;padding: 10px">
            <el-upload ref="field115" name="f" :file-list="field115fileList" action="http://127.0.0.1:8000/u">
                <el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
            </el-upload>
            <div v-for="(file,i) in files" :key="'file_'+i" class="item">
                <a :href="serverUrl+file">
                    <img :src="serverUrl+file"/>
                </a>
                <br/>
            </div>
        </div>
    </div>
</template>

<script>
    import request from "./utils/request"

    export default {
        name: 'App',
        data() {
            return {
                serverUrl: "http://localhost:8000/?p=",
                data: [],
                defaultProps: {
                    label: 'label',
                    children: 'children',
                },
                thisPath: null,
                files: [],
                field115fileList:[],
            }
        },
        mounted() {
            request.get('/dirserv/', {params: {p: "/"}}).then(res => {
                this.files = res.data.files
            })
        },
        methods: {
            handleNodeClick(data, resolve) {
                console.log(data.label)
                this.thisPath = data

                request.get('/dirserv/', {params: {p: data.label}}).then(res => {
                    let nodes = []
                    res.data.dirs.map(dir => {
                        nodes.push({label: dir, children: []})
                    })
                    resolve(nodes)
                })
            },
            handleNodeClick1(data) {
                console.log(data.label)
                this.thisPath = data
                this.files = []

                request.get('/dirserv/', {params: {p: data.label}}).then(res => {
                    this.files = res.data.files
                })
            },
        }
    }
</script>

<style>
    body {
        margin: 0;
        padding: 0;
    }

    .left {
        border-right: 1px solid #ff0000
    }

    .item {
        margin: 5px 0 0 5px;
    }

    .item a {
    }
</style>
