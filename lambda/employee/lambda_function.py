import json
import os

from bson import ObjectId
from pymongo import MongoClient


mongodb_uri = os.environ["MONGODB_URI"]
database_name = os.environ["MONGODB_DATABASE"]

client = MongoClient(
    mongodb_uri,
    serverSelectionTimeoutMS=5000
)

db = client[database_name]
employees_collection = db["employees"]


def response(status_code, body):
    return {
        "statusCode": status_code,
        "headers": {
            "Content-Type": "application/json"
        },
        "body": json.dumps(body)
    }


def lambda_handler(event, context):

    try:

        method = event.get("requestContext", {}).get("http", {}).get("method")

        path_parameters = event.get("pathParameters") or {}
        employee_id = path_parameters.get("id")

        # GET /employees/search
        if method == "GET" and event.get("rawPath") == "/employees/search":
            query_parameters = event.get("queryStringParameters") or {}
            query = query_parameters.get("q", "").strip()

            if not query:
                return response(400, {"error": "search query is required"})

            regex = {"$regex": query, "$options": "i"}

            employees = list(
                employees_collection.find({
                    "$or": [
                        {"name": regex},
                        {"email": regex},
                        {"department": regex},
                        {"position": regex}
                    ]
                })
            )

            for employee in employees:
                employee["_id"] = str(employee["_id"])

            return response(
                200,
                {
                    "employees": employees
                }
            )

        # --------------------------------
        # GET /employees/{id}
        # --------------------------------
        if method == "GET" and employee_id:

            if not ObjectId.is_valid(employee_id):
                return response(
                    400,
                    {
                        "error": "invalid employee id"
                    }
                )

            employee = employees_collection.find_one(
                {
                    "_id": ObjectId(employee_id)
                }
            )

            if employee is None:
                return response(
                    404,
                    {
                        "error": "employee not found"
                    }
                )

            employee["_id"] = str(employee["_id"])

            return response(
                200,
                {
                    "employee": employee
                }
            )

        # --------------------------------
        # GET /employees
        # --------------------------------
        if method == "GET":

            employees = list(
                employees_collection.find({})
            )

            for employee in employees:
                employee["_id"] = str(employee["_id"])

            return response(
                200,
                {
                    "employees": employees
                }
            )

        # --------------------------------
        # POST /employees
        # --------------------------------
        if method == "POST":

            body = event.get("body")

            if body:
                employee = json.loads(body)
            else:
                employee = event

            name = employee.get("name")
            email = employee.get("email")
            department = employee.get("department")
            position = employee.get("position")
            salary = employee.get("salary")

            if not name or not email or not department or not position or not salary:
                return response(
                    400,
                    {
                        "error": "name, email , department , position and salary are required"
                    }
                )

            result = employees_collection.insert_one({
                "name": name,
                "email": email,
                "department": department,
                "position": position,
                "salary": salary
            })

            return response(
                201,
                {
                    "message": "Employee created successfully",
                    "employee_id": str(result.inserted_id)
                }
            )
        
        # --------------------------------
        # updatee /employees/{id}
        # --------------------------------
        if method == "PUT" and employee_id:
            if not ObjectId.is_valid(employee_id):
                return response(
                    400,
                    {
                        "error": "invalid employee id"
                    }
                )
            
            body = event.get("body")

            if not body:
                return response(
                    400,
                    {
                        "error": "request body is required"
                    }
                )
            
            employee_data = json.loads(body)

            name = employee_data.get("name")
            email = employee_data.get("email")
            department = employee_data.get("department")
            position = employee_data.get("position")
            salary = employee_data.get("salary")

            if not name or not email or not department or not position or not salary:
                return response(
                    400,
                    {
                        "error": "name, email , department , position and salary are required"
                    }
                )

            result = employees_collection.update_one(
                {"_id": ObjectId(employee_id)},
                {
                    "$set": {
                        "name": name,
                        "email": email,
                        "department": department,
                        "position": position,
                        "salary": salary
                    }
                }
            )

            if result.matched_count == 0:
                return response(
                    404,
                    {
                        "error": "employee not found"
                    }
                )

            updated_employee = employees_collection.find_one(
                {"_id": ObjectId(employee_id)}
            )

            if not updated_employee:
                return {
                    "statusCode": 404,
                    "headers": {
                        "Content-Type": "application/json"
                    },
                    "body": json.dumps({
                        "error": "Employee not found"
                    })
                }

            updated_employee["_id"] = str(updated_employee["_id"])

            return {
                "statusCode": 200,
                "headers": {
                    "Content-Type": "application/json"
                },
                "body": json.dumps(updated_employee)
            }
        
        # --------------------------------
        # DELETE /employees/{id}
        # --------------------------------
        if method == "DELETE" and employee_id:

            if not ObjectId.is_valid(employee_id):
                return response(
                    400,
                    {
                        "error": "invalid employee id"
                    }
                )

            result = employees_collection.delete_one(
                {
                    "_id": ObjectId(employee_id)
                }
            )

            if result.deleted_count == 0:
                return response(
                    404,
                    {
                        "error": "employee not found"
                    }
                )

            return response(
                200,
                {
                    "message": "Employee deleted successfully",
                    "employee_id": employee_id
                }
            )
    
    

    except Exception as e:

        print("Error:")
        print(str(e))

        return response(
            500,
            {
                "error": "Internal server error"
            }
        )