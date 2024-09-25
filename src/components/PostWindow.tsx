import './css/PostWindow.css';
import sample_api_data from '../assets/sample_api_data';
import Post from './Post';

function PostWindow() {
    return (
        <div className="postWindowContainer">
            {sample_api_data.map((item) => {
               return <Post key={item.ID} post={item} />;
            })}
        </div>
    );
}

export default PostWindow;

