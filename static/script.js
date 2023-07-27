<script src="/static/js/script.js"></script>
const taskForm = document.getElementById('task-form');
const taskInput = document.getElementById('task-input');
const taskList = document.getElementById('task-list');

// Function to create a task item
function createTaskItem(task) {
    const li = document.createElement('li');
    li.innerHTML = `<span data-id="${task.id}" class="delete-task">X</span>${task.title}`;
    taskList.appendChild(li);
}

// Function to fetch tasks from the server
async function fetchTasks() {
    const response = await fetch('/tasks');
    const tasks = await response.json();
    tasks.forEach(task => {
        createTaskItem(task);
    });
}

// Function to handle task submission
async function handleTaskSubmit(event) {
    event.preventDefault();
    const taskTitle = taskInput.value.trim();

    if (taskTitle === '') {
        return;
    }

    const response = await fetch('/tasks/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title: taskTitle })
    });

    if (response.ok) {
        const task = await response.json();
        createTaskItem(task);
        taskInput.value = '';
    }
}

// Function to handle task deletion
async function handleTaskDeletion(event) {
    if (event.target.classList.contains('delete-task')) {
        const taskId = event.target.dataset.id;

        const response = await fetch('/tasks/delete', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ id: taskId })
        });

        if (response.ok) {
            event.target.parentElement.remove();
        }
    }
}

// Event listeners
taskForm.addEventListener('submit', handleTaskSubmit);
taskList.addEventListener('click', handleTaskDeletion);

// Fetch tasks on page load
fetchTasks();

