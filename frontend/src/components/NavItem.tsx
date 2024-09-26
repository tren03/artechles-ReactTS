import { Link } from 'react-router-dom';

type NavItemAttr = {
    to : "/"| "/posts"|"/home"|"/contact"|string
    label: string;
};

function NavItem({to , label}:NavItemAttr) {
    return (
        <Link to={to} className="nav-link">
            {label}
        </Link>
    );
}

export default NavItem;
