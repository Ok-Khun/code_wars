def better_than_average(class_points, your_points):
    # calculate average score
    avg_points = sum(class_points) / len(class_points)
    # compre with your points
    return your_points > avg_points 