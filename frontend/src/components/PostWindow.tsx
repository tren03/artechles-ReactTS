import './css/PostWindow.css';
import Post from './Post';
import PostContext from '../context/PostContext';
import { useContext, useEffect, useState } from 'react';
import ApiData from '../global_types/apiData';

function PostWindow() {
    // do not need setreqcategory function here
    const { apiData, requiredCategory } = useContext(PostContext);
    const [fileteredApiData,setFilteredApiData] = useState<ApiData[]>([])
    console.log(requiredCategory, apiData);
    console.log('req category changed' + requiredCategory);

    useEffect(()=>{
    
        

        let temp : ApiData[]
        temp = []
        if(requiredCategory==="all") {
            setFilteredApiData(apiData)
        }else{
        temp = []
        apiData.map((singlePost)=>{
            if(singlePost.Category==requiredCategory){
                temp.push(singlePost)
            }
        })
        setFilteredApiData(temp)    
        
        }
        
        console.log("this is the length of temp = "+temp.length)
         

    },[apiData,requiredCategory])


    return (
        <div className="postWindowContainer">
            {fileteredApiData.map((item) => {
                return <Post key={item.ID} post={item} />;
            })}
        </div>
    );
}

export default PostWindow;
