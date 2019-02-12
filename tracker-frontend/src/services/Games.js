import axios from 'axios';

export default {
    getGames() {
        return axios.get("/games")
            .then(response => {
                return response.data;
            });
    }
}
