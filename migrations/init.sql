create table ingredients (
    id   uuid primary key,
    name varchar not null
);

create table foods (
    id               uuid primary key,
    name             varchar not null,
    proteins         double precision,
    fats             double precision,
    carbs            double precision,
    calories         double precision,
    cooking_duration bigint
);

create table categories (
    id   uuid primary key,
    name varchar not null
);

create table meals (
    id   uuid primary key,
    name varchar not null
);

create table days (
    id   uuid primary key,
    date date not null
);

create table ingredient_food (
    ingredient_id uuid references ingredients(id) on delete cascade,
    food_id       uuid references foods(id) on delete cascade,
    primary key (ingredient_id, food_id)
);

create table food_category (
    food_id     uuid references foods(id) on delete cascade,
    category_id uuid references categories(id) on delete cascade,
    primary key (food_id, category_id)
);

create table food_meal (
    food_id uuid references foods(id) on delete cascade,
    meal_id uuid references meals(id) on delete cascade,
    primary key (food_id, meal_id)
);

create table meal_day (
    meal_id uuid references meals(id) on delete cascade,
    day_id  uuid references foods(id) on delete cascade,
    primary key (meal_id, day_id)
);