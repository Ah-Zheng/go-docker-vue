<template lang="pug">
    .home
        button(@click="getNumber") 取得數字
        p
        input(
            type="text"
            v-model="text"
        )
        button(@click="guest" style="margin-left: 10px") 送出
        p(
            v-for="(data, index) in datas"
            :key="index"
        ) {{ data.send }} {{ data.res }}
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import axios from 'axios';

@Component
export default class Home extends Vue {
    private text = '';

    private datas: any[] = [];

    private getNumber(): void {
        axios.get('/api/generate').then((res) => {
            console.log('res :>> ', res);
        });
    }

    private guest() {
        if (this.text) {
            axios.post('/api/guest', { number: this.text }).then((res) => {
                this.datas.push({
                    send: this.text,
                    res: res.data.msg
                });
            });
        }
    }
}
</script>
