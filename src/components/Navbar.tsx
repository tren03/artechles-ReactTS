import './css/Navbar.css';
import NavItem from './NavItem';

const Navbar = () => {

    return (
        <nav className="navbar">

            <ul className='nav-container'
            >
                <li className="nav-item">
                    <NavItem
                        to="/"
                        label="home"
                    />
                </li>
                <li className="nav-item">
                    <NavItem
                        to="/posts"
                        label="posts"
                    />
                </li>
                <li className="nav-item">
                    <NavItem
                        to="/about"
                        label="about"
                    />
                </li>
                <li className="nav-item">
                    <NavItem
                        to="/contact"
                        label="contact"
                    />
                </li>
            </ul>
        </nav>
    );
};

export default Navbar;
