import boto3

client = boto3.client('ce', region_name='us-east-1')

response = client.get_cost_and_usage(
    TimePeriod={
        'Start': '2022-07-01',
        'End': '2022-07-31'
    },
    Granularity='MONTHLY',
    Metrics=[
        'AmortizedCost',
    ]
)

print(response)
