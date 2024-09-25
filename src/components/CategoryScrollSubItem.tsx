type TitleProps = {
    title: string | null;
}; 

function CategoryScrollSubItem({ title }: TitleProps) {
    function handleOnClickCategory(){
        console.log("clicked")
    }

    return <div onClick={handleOnClickCategory}> {title} </div>;
}

export default CategoryScrollSubItem;

