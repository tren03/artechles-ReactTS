import { useEffect, useState } from 'react';
import PostWindow from '../components/PostWindow';
import ScrollWindow from '../components/ScrollWindow';
import PostContext from '../context/PostContext';
import './css/home.css';
import sample_api_data from '../assets/sample_api_data';
import ApiData from '../global_types/apiData';
import axios from 'axios';
const Home = () => {
    const [apiData, setApidata] = useState<ApiData[]>([]);
    const [requiredCategory, setRequiredCategory] = useState<string>('all');

    useEffect(() => {
        setApidata(sample_api_data); // Initially set sample data

        async function getAllPosts() {
            try {
                // Fetch the data from the API
                const response = await axios.get<ApiData[]>('http://localhost:3000/getallposts');

                // Log the response data
                console.log(response.data);

                // Set the response data (array of ApiData) into your state
                setApidata(response.data);
            } catch (error) {
                console.log('Error fetching posts:', error);
            }
        }

        // Fetch the posts when the component mounts
        getAllPosts();
    }, []);
    return (
        <PostContext.Provider
            value={{
                apiData,
                requiredCategory,
                setRequiredCategory,
            }}
        >
            <div className="home-container">
                <div className="scrollWindowSection">
                    <ScrollWindow />
                </div>
                <div className="postWindowSection">
                    <PostWindow />
                </div>
            </div>
        </PostContext.Provider>
    );
};

export default Home;
