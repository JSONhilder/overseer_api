CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    u_id INTEGER NOT NULL,
    project_name VARCHAR (100) NOT NULL,
    project_desc VARCHAR (255) NOT NULL,
    project_time interval,
    project_completed BOOLEAN DEFAULT '0' NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    p_id INTEGER NOT NULL,
    task_name VARCHAR (100) NOT NULL,
    task_desc VARCHAR (255) NOT NULL,
    task_time interval,
    task_completed BOOLEAN DEFAULT '0' NOT NULL,
    constraint fk_project_tasks
        foreign key (p_id) 
            REFERENCES projects (id)
            ON DELETE CASCADE
);