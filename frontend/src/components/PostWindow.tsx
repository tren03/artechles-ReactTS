import './css/PostWindow.css';
import Post from './Post';
import PostContext from '../context/PostContext';
import { useContext, useEffect, useState } from 'react';
import ApiData from '../global_types/apiData';
import SearchFilter from './SearchFilter';

function PostWindow() {
    // do not need setreqcategory function here
    const { apiData, requiredCategory } = useContext(PostContext);
    const [filteredApiData, setFilteredApiData] = useState<ApiData[]>([]);
    // needed to restore all posts after fuzzy search return no posts
    const [unfilteredApiData,setUnfilteredApiData] = useState<ApiData[]>([]);
    console.log(requiredCategory, apiData);
    console.log('req category changed' + requiredCategory);

    useEffect(() => {
        let temp_unfiltered: ApiData[];
        temp_unfiltered = [];
        if (requiredCategory === 'all') {
            setUnfilteredApiData(apiData);
        } else {
            temp_unfiltered = [];
            apiData.map((singlePost) => {
                if (singlePost.Category == requiredCategory) {
                    temp_unfiltered.push(singlePost);
                }
            });
            setUnfilteredApiData(temp_unfiltered);
        }

        let temp: ApiData[];
        temp = [];
        if (requiredCategory === 'all') {
            setFilteredApiData(apiData);
        } else {
            temp = [];
            apiData.map((singlePost) => {
                if (singlePost.Category == requiredCategory) {
                    temp.push(singlePost);
                }
            });
            setFilteredApiData(temp);
        }

        console.log('this is the length of temp = ' + temp.length);
    }, [apiData, requiredCategory]);

    return (
        <div className="postWindowContainer">
            <SearchFilter filteredApiData={filteredApiData} setFilteredApiData={setFilteredApiData} unfilteredApiData={unfilteredApiData}/>
            {filteredApiData.map((item) => {
                return <Post key={item.ID} post={item} />;
            })}
        </div>
    );
}

export default PostWindow;
