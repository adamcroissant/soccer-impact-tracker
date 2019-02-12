import axios from 'axios';

export default {
    getGames() {
        return axios.get("/getgames1")
            .then(response => {
                return response.data;
            });
    }
}
