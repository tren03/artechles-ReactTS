import './App.css';
import Navbar from './components/Navbar';
import {
    BrowserRouter as Router,
    Route,
    Routes,
} from 'react-router-dom';

import Home from './pages/Home';
import Posts from './pages/Posts';
import About from './pages/About';
import Contact from './pages/Contact';

function App() {
    return (
        <div className="app-container">
            <Router>
                <Navbar />
                <Routes>
                    <Route
                        path="/"
                        element={<Home />}
                    />
                    <Route
                        path="/posts"
                        element={<Posts />}
                    />
                    <Route
                        path="/about"
                        element={<About />}
                    />
                    <Route
                        path="/contact"
                        element={<Contact />}
                    />
                </Routes>
            </Router>
        </div>
    );
}

export default App;
