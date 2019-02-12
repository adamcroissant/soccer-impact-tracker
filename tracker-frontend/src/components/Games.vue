<template>
    <div>
        <button id="loadGames" v-on:click="load">Load Game History</button>
        <table id="games">
            <tr>
                <th>Opponent</th>
                <th>Date</th>
                <th>Result</th>
                <th>Attended?</th>
            </tr>
            <tr v-for="row in rows" :key="row.id">
                <td :id="'game-' + row.id">{{row.Opponent}}</td>
                <td>{{row.Date}}</td>
                <td>{{row.Result}}</td>
                <td>{{row.DidAttend}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
import Games from '@/services/Games'

export default {
    name: "directory",
    data() {
        return {
            rows: []
        }
    },
    methods: {
        load: function() {
            Games.getGames().then(games => {
                this.rows = games;
            });
        }
    }
}
</script>

<style scoped>
button {
    margin-bottom: 10px;
}
table {
    width: 400px;
    margin: 0 auto;
}
td {
    padding: 10px;
}
</style>
