import "./css/categoryScrollSubItem.css"
type TitleProps = {
    
    title: string | null;
}; 

function CategoryScrollSubItem({ title }: TitleProps) {
    function handleOnClickCategory(){
        console.log("clicked")
    }

    return <button className="CategoryScrollSubItem" onClick={handleOnClickCategory}> {title} </button >;
}

export default CategoryScrollSubItem;

