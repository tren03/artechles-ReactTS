import { useState } from 'react';
import './css/categoryScrollItem.css';
import CategoryScrollSubItem from './CategoryScrollSubItem';
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

function CategoryScrollItem({
    categoryName,
    posts,
}: CategoryItem) {
    const [isDropDown, setDropDown] =
        useState(false);
    // Event handler for button click

    let SubItems: JSX.Element | null;

    if (isDropDown) {
        SubItems = (
            <div className='dropDownContainer'>
                {posts.map((post) => {
                    return (
                        <CategoryScrollSubItem
                            key={post.ID}
                            title={post.Title}
                        />
                    );
                })}
            </div>
        );
    }
    function handleOnClickCategory() {
        console.log('clicked');
        setDropDown(!isDropDown);
    }

    return (
        <div className='categoryItemContainer'>
            <button
                className="CategoryItem"
                onClick={handleOnClickCategory}
            >
                {categoryName}
            </button>
            {SubItems}
        </div>
    );
}

export default CategoryScrollItem;
