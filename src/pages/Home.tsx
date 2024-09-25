import PostWindow from '../components/PostWindow';
import ScrollWindow from '../components/ScrollWindow';
import './css/home.css';
const Home = () => {
    return (
        <div className="home-container">
            <div className="scrollWindowSection">
                <ScrollWindow />
            </div>
            <div className="postWindowSection">
                <PostWindow />
            </div>
        </div>
    );
};

export default Home;
