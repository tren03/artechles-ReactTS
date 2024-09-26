import { useEffect, useState } from 'react';
import PostWindow from '../components/PostWindow';
import ScrollWindow from '../components/ScrollWindow';
import PostContext from '../context/PostContext';
import './css/home.css';
import sample_api_data from '../assets/sample_api_data';
import ApiData from '../global_types/apiData';
const Home = () => {
    const [apiData, setApidata] = useState<ApiData[]>([]);
    const [requiredCategory, setRequiredCategory] = useState<string>('all');

    useEffect(() => {
        setApidata(sample_api_data);
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
