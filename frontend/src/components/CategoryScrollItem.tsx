import { useContext, useState } from 'react';
import './css/categoryScrollItem.css';
import CategoryScrollSubItem from './CategoryScrollSubItem';
import PostContext from '../context/PostContext';
type Post = {
    ID: number;
    Title: string;
    Url: string;
    Body: string;
    Date: string;
};

type CategoryItem = {
    categoryName: string;
    posts: Post[];
};

function CategoryScrollItem({ categoryName, posts }: CategoryItem) {
    const [isDropDown, setDropDown] = useState(false);
    // Event handler for button click
        
    const {setRequiredCategory} = useContext(PostContext)

    let SubItems: JSX.Element | null;
    SubItems = null

    if (isDropDown) {
        SubItems = (
            <div className="dropDownContainer">
                {posts.map((post) => {
                    return <CategoryScrollSubItem key={post.ID} title={post.Title} />;
                })}
            </div>
        );
    }
    function handleOnClickCategory(categoryName:string) {
        console.log('clicked the category '+categoryName);
        setRequiredCategory(categoryName)
        setDropDown(!isDropDown);
    }

    return (
        <div className="categoryItemContainer">
            <button className="CategoryItem" onClick={() =>handleOnClickCategory(categoryName)}>
                {categoryName}
            </button>
            {SubItems}
        </div>
    );
}

export default CategoryScrollItem;
